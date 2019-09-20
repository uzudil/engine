// Code generated by G3NSHADERS. DO NOT EDIT.
// To regenerate this file install 'g3nshaders' and execute:
// 'go generate' in this folder.

package shaders

const include_attributes_source = `//
// Vertex attributes
//
layout(location = 0) in  vec3  VertexPosition;
layout(location = 1) in  vec3  VertexNormal;
layout(location = 2) in  vec3  VertexColor;
layout(location = 3) in  vec2  VertexTexcoord;
`

const include_bones_vertex_source = `#ifdef BONE_INFLUENCERS
    #if BONE_INFLUENCERS > 0

        mat4 influence = mBones[int(matricesIndices[0])] * matricesWeights[0];
        #if BONE_INFLUENCERS > 1
            influence += mBones[int(matricesIndices[1])] * matricesWeights[1];
            #if BONE_INFLUENCERS > 2
                influence += mBones[int(matricesIndices[2])] * matricesWeights[2];
                #if BONE_INFLUENCERS > 3
                    influence += mBones[int(matricesIndices[3])] * matricesWeights[3];
    //                #if BONE_INFLUENCERS > 4
    //                    influence += mBones[int(matricesIndicesExtra[0])] * matricesWeightsExtra[0];
    //                    #if BONE_INFLUENCERS > 5
    //                        influence += mBones[int(matricesIndicesExtra[1])] * matricesWeightsExtra[1];
    //                        #if BONE_INFLUENCERS > 6
    //                            influence += mBones[int(matricesIndicesExtra[2])] * matricesWeightsExtra[2];
    //                            #if BONE_INFLUENCERS > 7
    //                                influence += mBones[int(matricesIndicesExtra[3])] * matricesWeightsExtra[3];
    //                            #endif
    //                        #endif
    //                    #endif
    //                #endif
                #endif
            #endif
        #endif

        finalWorld = finalWorld * influence;

    #endif
#endif
`

const include_bones_vertex_declaration_source = `#ifdef BONE_INFLUENCERS
    #if BONE_INFLUENCERS > 0
	uniform mat4 mBones[TOTAL_BONES];
    in vec4 matricesIndices;
    in vec4 matricesWeights;
//    #if BONE_INFLUENCERS > 4
//        in vec4 matricesIndicesExtra;
//        in vec4 matricesWeightsExtra;
//    #endif
    #endif
#endif
`

const include_lights_source = `//
// Lights uniforms
//

#if AMB_LIGHTS>0
    // Ambient lights color uniform
    uniform vec3 AmbientLightColor[AMB_LIGHTS];
#endif

#if DIR_LIGHTS>0
    // Directional lights uniform array. Each directional light uses 2 elements
    uniform vec3 DirLight[2*DIR_LIGHTS];
    // Macros to access elements inside the DirectionalLight uniform array
    #define DirLightColor(a)		DirLight[2*a]
    #define DirLightPosition(a)		DirLight[2*a+1]
#endif

#if POINT_LIGHTS>0
    // Point lights uniform array. Each point light uses 3 elements
    uniform vec3 PointLight[3*POINT_LIGHTS];
    // Macros to access elements inside the PointLight uniform array
    #define PointLightColor(a)			PointLight[3*a]
    #define PointLightPosition(a)		PointLight[3*a+1]
    #define PointLightLinearDecay(a)	PointLight[3*a+2].x
    #define PointLightQuadraticDecay(a)	PointLight[3*a+2].y
#endif

#if SPOT_LIGHTS>0
    // Spot lights uniforms. Each spot light uses 5 elements
    uniform vec3  SpotLight[5*SPOT_LIGHTS];
    // Macros to access elements inside the PointLight uniform array
    #define SpotLightColor(a)			SpotLight[5*a]
    #define SpotLightPosition(a)		SpotLight[5*a+1]
    #define SpotLightDirection(a)		SpotLight[5*a+2]
    #define SpotLightAngularDecay(a)	SpotLight[5*a+3].x
    #define SpotLightCutoffAngle(a)		SpotLight[5*a+3].y
    #define SpotLightLinearDecay(a)		SpotLight[5*a+3].z
    #define SpotLightQuadraticDecay(a)	SpotLight[5*a+4].x
#endif
`

const include_material_source = `//
// Material properties uniform
//

// Material parameters uniform array
uniform vec3 Material[6];
// Macros to access elements inside the Material array
#define MatAmbientColor		Material[0]
#define MatDiffuseColor     Material[1]
#define MatSpecularColor    Material[2]
#define MatEmissiveColor    Material[3]
#define MatShininess        Material[4].x
#define MatOpacity          Material[4].y
#define MatPointSize        Material[4].z
#define MatPointRotationZ   Material[5].x

#if MAT_TEXTURES > 0
    // Texture unit sampler array
    uniform sampler2D MatTexture[MAT_TEXTURES];
    // Texture parameters (3*vec2 per texture)
    uniform vec2 MatTexinfo[3*MAT_TEXTURES];
    // Macros to access elements inside the MatTexinfo array
    #define MatTexOffset(a)		MatTexinfo[(3*a)]
    #define MatTexRepeat(a)		MatTexinfo[(3*a)+1]
    #define MatTexFlipY(a)		bool(MatTexinfo[(3*a)+2].x)
    #define MatTexVisible(a)	bool(MatTexinfo[(3*a)+2].y)
    // Alpha compositing (see here: https://ciechanow.ski/alpha-compositing/)
    vec4 Blend(vec4 texMixed, vec4 texColor) {
        texMixed.rgb *= texMixed.a;
        texColor.rgb *= texColor.a;
        texMixed = texColor + texMixed * (1 - texColor.a);
        if (texMixed.a > 0.0) {
            texMixed.rgb /= texMixed.a;
        }
        return texMixed;
    }
#endif
`

const include_morphtarget_vertex_source = `#ifdef MORPHTARGETS

    #include <morphtarget_vertex2> [MORPHTARGETS]

#endif
`

const include_morphtarget_vertex2_source = `	vPosition += MorphPosition{i} * morphTargetInfluences[{i}];
  #ifdef MORPHTARGETS_NORMAL
	vNormal += MorphNormal{i} * morphTargetInfluences[{i}];
  #endif`

const include_morphtarget_vertex_declaration_source = `#ifdef MORPHTARGETS
	uniform float morphTargetInfluences[MORPHTARGETS];
	#include <morphtarget_vertex_declaration2> [MORPHTARGETS]
#endif
`

const include_morphtarget_vertex_declaration2_source = `	in vec3 MorphPosition{i};
  #ifdef MORPHTARGETS_NORMAL
	in vec3 MorphNormal{i};
  #endif
`

const include_phong_model_source = `/***
 phong lighting model
 Parameters:
    position:   input vertex position in camera coordinates
    normal:     input vertex normal in camera coordinates
    camDir:     input camera directions
    matAmbient: input material ambient color
    matDiffuse: input material diffuse color
    ambdiff:    output ambient+diffuse color
    spec:       output specular color
 Uniforms:
    AmbientLightColor[]
    DiffuseLightColor[]
    DiffuseLightPosition[]
    PointLightColor[]
    PointLightPosition[]
    PointLightLinearDecay[]
    PointLightQuadraticDecay[]
    MatSpecularColor
    MatShininess
*****/
void phongModel(vec4 position, vec3 normal, vec3 camDir, vec3 matAmbient, vec3 matDiffuse, out vec3 ambdiff, out vec3 spec) {

    vec3 ambientTotal  = vec3(0.0);
    vec3 diffuseTotal  = vec3(0.0);
    vec3 specularTotal = vec3(0.0);

    bool noLights = true;
    const float EPS = 0.00001;

#if AMB_LIGHTS>0
    noLights = false;
    // Ambient lights
    for (int i = 0; i < AMB_LIGHTS; ++i) {
        ambientTotal += AmbientLightColor[i] * matAmbient;
    }
#endif

#if DIR_LIGHTS>0
    noLights = false;
    // Directional lights
    for (int i = 0; i < DIR_LIGHTS; ++i) {
        vec3 lightDirection = normalize(DirLightPosition(i)); // Vector from fragment to light source
        float dotNormal = dot(lightDirection, normal); // Dot product between light direction and fragment normal // TODO can remove the max here
        if (dotNormal > EPS) { // If the fragment is lit
            diffuseTotal += DirLightColor(i) * matDiffuse * dotNormal;
            specularTotal += DirLightColor(i) * MatSpecularColor * pow(max(dot(reflect(-lightDirection, normal), camDir), 0.0), MatShininess);
        }
    }
#endif

#if POINT_LIGHTS>0
    noLights = false;
    // Point lights
    for (int i = 0; i < POINT_LIGHTS; ++i) {
        vec3 lightDirection = PointLightPosition(i) - vec3(position); // Vector from fragment to light source
        float lightDistance = length(lightDirection); // Distance from fragment to light source
        lightDirection = lightDirection / lightDistance; // Normalize lightDirection
        float dotNormal = dot(lightDirection, normal);  // Dot product between light direction and fragment normal
        if (dotNormal > EPS) { // If the fragment is lit
            float attenuation = 1.0 / (1.0 + PointLightLinearDecay(i) * lightDistance + PointLightQuadraticDecay(i) * lightDistance * lightDistance);
            vec3 attenuatedColor = PointLightColor(i) * attenuation;
            diffuseTotal += attenuatedColor * matDiffuse * dotNormal;
            specularTotal += attenuatedColor * MatSpecularColor * pow(max(dot(reflect(-lightDirection, normal), camDir), 0.0), MatShininess);
        }
    }
#endif

#if SPOT_LIGHTS>0
    noLights = false;
    for (int i = 0; i < SPOT_LIGHTS; ++i) {
        // Calculates the direction and distance from the current vertex to this spot light.
        vec3 lightDirection = SpotLightPosition(i) - vec3(position); // Vector from fragment to light source
        float lightDistance = length(lightDirection); // Distance from fragment to light source
        lightDirection = lightDirection / lightDistance; // Normalize lightDirection
        float angleDot = dot(-lightDirection, SpotLightDirection(i));
        float angle = acos(angleDot);
        float cutoff = radians(clamp(SpotLightCutoffAngle(i), 0.0, 90.0));
        if (angle < cutoff) { // Check if fragment is inside spotlight beam
            float dotNormal = dot(lightDirection, normal); // Dot product between light direction and fragment normal
            if (dotNormal > EPS) { // If the fragment is lit
                float attenuation = 1.0 / (1.0 + SpotLightLinearDecay(i) * lightDistance + SpotLightQuadraticDecay(i) * lightDistance * lightDistance);
                float spotFactor = pow(angleDot, SpotLightAngularDecay(i));
                vec3 attenuatedColor = SpotLightColor(i) * attenuation * spotFactor;
                diffuseTotal += attenuatedColor * matDiffuse * dotNormal;
                specularTotal += attenuatedColor * MatSpecularColor * pow(max(dot(reflect(-lightDirection, normal), camDir), 0.0), MatShininess);
            }
        }
    }
#endif
    if (noLights) {
        diffuseTotal = matDiffuse;
    }
    // Sets output colors
    ambdiff = ambientTotal + MatEmissiveColor + diffuseTotal;
    spec = specularTotal;
}
`

const basic_fragment_source = `precision highp float;

in vec3 Color;
out vec4 FragColor;

void main() {

    FragColor = vec4(Color, 1.0);
}
`

const basic_vertex_source = `#include <attributes>

// Model uniforms
uniform mat4 MVP;

// Final output color for fragment shader
out vec3 Color;

void main() {

    Color = VertexColor;
    gl_Position = MVP * vec4(VertexPosition, 1.0);
}
`

const panel_fragment_source = `precision highp float;

// Texture uniforms
uniform sampler2D	MatTexture;
uniform vec2		MatTexinfo[3];

// Macros to access elements inside the MatTexinfo array
#define MatTexOffset		MatTexinfo[0]
#define MatTexRepeat		MatTexinfo[1]
#define MatTexFlipY	    	bool(MatTexinfo[2].x) // not used
#define MatTexVisible	    bool(MatTexinfo[2].y) // not used

// Inputs from vertex shader
in vec2 FragTexcoord;

// Input uniform
uniform vec4 Panel[8];
#define Bounds			Panel[0]		  // panel bounds in texture coordinates
#define Border			Panel[1]		  // panel border in texture coordinates
#define Padding			Panel[2]		  // panel padding in texture coordinates
#define Content			Panel[3]		  // panel content area in texture coordinates
#define BorderColor		Panel[4]		  // panel border color
#define PaddingColor	Panel[5]		  // panel padding color
#define ContentColor	Panel[6]		  // panel content color
#define TextureValid	bool(Panel[7].x)  // texture valid flag

// Output
out vec4 FragColor;


/***
* Checks if current fragment texture coordinate is inside the
* supplied rectangle in texture coordinates:
* rect[0] - position x [0,1]
* rect[1] - position y [0,1]
* rect[2] - width [0,1]
* rect[3] - height [0,1]
*/
bool checkRect(vec4 rect) {

    if (FragTexcoord.x < rect[0]) {
        return false;
    }
    if (FragTexcoord.x > rect[0] + rect[2]) {
        return false;
    }
    if (FragTexcoord.y < rect[1]) {
        return false;
    }
    if (FragTexcoord.y > rect[1] + rect[3]) {
        return false;
    }
    return true;
}


void main() {

    // Discard fragment outside of received bounds
    // Bounds[0] - xmin
    // Bounds[1] - ymin
    // Bounds[2] - xmax
    // Bounds[3] - ymax
    if (FragTexcoord.x <= Bounds[0] || FragTexcoord.x >= Bounds[2]) {
        discard;
    }
    if (FragTexcoord.y <= Bounds[1] || FragTexcoord.y >= Bounds[3]) {
        discard;
    }

    // Check if fragment is inside content area
    if (checkRect(Content)) {

        // If no texture, the color will be the material color.
        vec4 color = ContentColor;

		if (TextureValid) {
            // Adjust texture coordinates to fit texture inside the content area
            vec2 offset = vec2(-Content[0], -Content[1]);
            vec2 factor = vec2(1.0/Content[2], 1.0/Content[3]);
            vec2 texcoord = (FragTexcoord + offset) * factor;
            vec4 texColor = texture(MatTexture, texcoord * MatTexRepeat + MatTexOffset);

            // Mix content color with texture color.
            // Note that doing a simple linear interpolation (e.g. using mix()) is not correct!
            // The right formula can be found here: https://en.wikipedia.org/wiki/Alpha_compositing#Alpha_blending
            // For a more in-depth discussion: http://apoorvaj.io/alpha-compositing-opengl-blending-and-premultiplied-alpha.html#toc4
            // Another great discussion here: https://ciechanow.ski/alpha-compositing/

            // Alpha premultiply the content color
            vec4 contentPre = ContentColor;
            contentPre.rgb *= contentPre.a;

            // Alpha premultiply the content color
            vec4 texPre = texColor;
            texPre.rgb *= texPre.a;

            // Combine colors to obtain the alpha premultiplied final color
            color = texPre + contentPre * (1.0 - texPre.a);

            // Un-alpha-premultiply
            color.rgb /= color.a;
		}

        FragColor = color;
        return;
    }

    // Checks if fragment is inside paddings area
    if (checkRect(Padding)) {
        FragColor = PaddingColor;
        return;
    }

    // Checks if fragment is inside borders area
    if (checkRect(Border)) {
        FragColor = BorderColor;
        return;
    }

    // Fragment is in margins area (always transparent)
    FragColor = vec4(1,1,1,0);
}
`

const panel_vertex_source = `#include <attributes>

// Model uniforms
uniform mat4 ModelMatrix;

// Outputs for fragment shader
out vec2 FragTexcoord;


void main() {

    // Always flip texture coordinates
    vec2 texcoord = VertexTexcoord;
    texcoord.y = 1.0 - texcoord.y;
    FragTexcoord = texcoord;

    // Set position
    vec4 pos = vec4(VertexPosition.xyz, 1);
    gl_Position = ModelMatrix * pos;
}
`

const physical_fragment_source = `//
// Physically Based Shading of a microfacet surface material - Fragment Shader
// Modified from reference implementation at https://github.com/KhronosGroup/glTF-WebGL-PBR
//
// References:
// [1] Real Shading in Unreal Engine 4
//     http://blog.selfshadow.com/publications/s2013-shading-course/karis/s2013_pbs_epic_notes_v2.pdf
// [2] Physically Based Shading at Disney
//     http://blog.selfshadow.com/publications/s2012-shading-course/burley/s2012_pbs_disney_brdf_notes_v3.pdf
// [3] README.md - Environment Maps
//     https://github.com/KhronosGroup/glTF-WebGL-PBR/#environment-maps
// [4] "An Inexpensive BRDF Model for Physically based Rendering" by Christophe Schlick
//     https://www.cs.virginia.edu/~jdl/bib/appearance/analytic%20models/schlick94b.pdf

//#extension GL_EXT_shader_texture_lod: enable
//#extension GL_OES_standard_derivatives : enable

precision highp float;

//uniform vec3 u_LightDirection;
//uniform vec3 u_LightColor;

//#ifdef USE_IBL
//uniform samplerCube u_DiffuseEnvSampler;
//uniform samplerCube u_SpecularEnvSampler;
//uniform sampler2D u_brdfLUT;
//#endif

#ifdef HAS_BASECOLORMAP
uniform sampler2D uBaseColorSampler;
#endif
#ifdef HAS_METALROUGHNESSMAP
uniform sampler2D uMetallicRoughnessSampler;
#endif
#ifdef HAS_NORMALMAP
uniform sampler2D uNormalSampler;
//uniform float uNormalScale;
#endif
#ifdef HAS_EMISSIVEMAP
uniform sampler2D uEmissiveSampler;
#endif
#ifdef HAS_OCCLUSIONMAP
uniform sampler2D uOcclusionSampler;
uniform float uOcclusionStrength;
#endif

// Material parameters uniform array
uniform vec4 Material[3];
// Macros to access elements inside the Material array
#define uBaseColor		    Material[0]
#define uEmissiveColor      Material[1]
#define uMetallicFactor     Material[2].x
#define uRoughnessFactor    Material[2].y

#include <lights>

// Inputs from vertex shader
in vec3 Position;       // Vertex position in camera coordinates.
in vec3 Normal;         // Vertex normal in camera coordinates.
in vec3 CamDir;         // Direction from vertex to camera
in vec2 FragTexcoord;

// Final fragment color
out vec4 FragColor;

// Encapsulate the various inputs used by the various functions in the shading equation
// We store values in this struct to simplify the integration of alternative implementations
// of the shading terms, outlined in the Readme.MD Appendix.
struct PBRLightInfo
{
    float NdotL;                  // cos angle between normal and light direction
    float NdotV;                  // cos angle between normal and view direction
    float NdotH;                  // cos angle between normal and half vector
    float LdotH;                  // cos angle between light direction and half vector
    float VdotH;                  // cos angle between view direction and half vector
};

struct PBRInfo
{
    float perceptualRoughness;    // roughness value, as authored by the model creator (input to shader)
    float metalness;              // metallic value at the surface
    vec3 reflectance0;            // full reflectance color (normal incidence angle)
    vec3 reflectance90;           // reflectance color at grazing angle
    float alphaRoughness;         // roughness mapped to a more linear change in the roughness (proposed by [2])
    vec3 diffuseColor;            // color contribution from diffuse lighting
    vec3 specularColor;           // color contribution from specular lighting
};

const float M_PI = 3.141592653589793;
const float c_MinRoughness = 0.04;

vec4 SRGBtoLINEAR(vec4 srgbIn) {
//#ifdef MANUAL_SRGB
//    #ifdef SRGB_FAST_APPROXIMATION
//        vec3 linOut = pow(srgbIn.xyz,vec3(2.2));
//    #else //SRGB_FAST_APPROXIMATION
        vec3 bLess = step(vec3(0.04045),srgbIn.xyz);
        vec3 linOut = mix( srgbIn.xyz/vec3(12.92), pow((srgbIn.xyz+vec3(0.055))/vec3(1.055),vec3(2.4)), bLess );
//    #endif //SRGB_FAST_APPROXIMATION
        return vec4(linOut,srgbIn.w);
//#else //MANUAL_SRGB
//    return srgbIn;
//#endif //MANUAL_SRGB
}

// Find the normal for this fragment, pulling either from a predefined normal map
// or from the interpolated mesh normal and tangent attributes.
vec3 getNormal()
{
    // Retrieve the tangent space matrix
//#ifndef HAS_TANGENTS
    vec3 pos_dx = dFdx(Position);
    vec3 pos_dy = dFdy(Position);
    vec3 tex_dx = dFdx(vec3(FragTexcoord, 0.0));
    vec3 tex_dy = dFdy(vec3(FragTexcoord, 0.0));
    vec3 t = (tex_dy.t * pos_dx - tex_dx.t * pos_dy) / (tex_dx.s * tex_dy.t - tex_dy.s * tex_dx.t);

//#ifdef HAS_NORMALS
    vec3 ng = normalize(Normal);
//#else
//    vec3 ng = cross(pos_dx, pos_dy);
//#endif

    t = normalize(t - ng * dot(ng, t));
    vec3 b = normalize(cross(ng, t));
    mat3 tbn = mat3(t, b, ng);
//#else // HAS_TANGENTS
//    mat3 tbn = v_TBN;
//#endif

#ifdef HAS_NORMALMAP
    float uNormalScale = 1.0;
    vec3 n = texture(uNormalSampler, FragTexcoord).rgb;
    n = normalize(tbn * ((2.0 * n - 1.0) * vec3(uNormalScale, uNormalScale, 1.0)));
#else
    // The tbn matrix is linearly interpolated, so we need to re-normalize
    vec3 n = normalize(tbn[2].xyz);
#endif

    return n;
}

// Calculation of the lighting contribution from an optional Image Based Light source.
// Precomputed Environment Maps are required uniform inputs and are computed as outlined in [1].
// See our README.md on Environment Maps [3] for additional discussion.
vec3 getIBLContribution(PBRInfo pbrInputs, PBRLightInfo pbrLight, vec3 n, vec3 reflection)
{
    float mipCount = 9.0; // resolution of 512x512
    float lod = (pbrInputs.perceptualRoughness * mipCount);
    // retrieve a scale and bias to F0. See [1], Figure 3
    vec3 brdf = vec3(0.5,0.5,0.5);//SRGBtoLINEAR(texture(u_brdfLUT, vec2(pbrLight.NdotV, 1.0 - pbrInputs.perceptualRoughness))).rgb;
    vec3 diffuseLight = vec3(0.5,0.5,0.5);//SRGBtoLINEAR(textureCube(u_DiffuseEnvSampler, n)).rgb;

//#ifdef USE_TEX_LOD
//    vec3 specularLight = SRGBtoLINEAR(textureCubeLodEXT(u_SpecularEnvSampler, reflection, lod)).rgb;
//#else
    vec3 specularLight = vec3(0.5,0.5,0.5);//SRGBtoLINEAR(textureCube(u_SpecularEnvSampler, reflection)).rgb;
//#endif

    vec3 diffuse = diffuseLight * pbrInputs.diffuseColor;
    vec3 specular = specularLight * (pbrInputs.specularColor * brdf.x + brdf.y);

    // For presentation, this allows us to disable IBL terms
//    diffuse *= u_ScaleIBLAmbient.x;
//    specular *= u_ScaleIBLAmbient.y;

    return diffuse + specular;
}

// Basic Lambertian diffuse
// Implementation from Lambert's Photometria https://archive.org/details/lambertsphotome00lambgoog
// See also [1], Equation 1
vec3 diffuse(PBRInfo pbrInputs)
{
    return pbrInputs.diffuseColor / M_PI;
}

// The following equation models the Fresnel reflectance term of the spec equation (aka F())
// Implementation of fresnel from [4], Equation 15
vec3 specularReflection(PBRInfo pbrInputs, PBRLightInfo pbrLight)
{
    return pbrInputs.reflectance0 + (pbrInputs.reflectance90 - pbrInputs.reflectance0) * pow(clamp(1.0 - pbrLight.VdotH, 0.0, 1.0), 5.0);
}

// This calculates the specular geometric attenuation (aka G()),
// where rougher material will reflect less light back to the viewer.
// This implementation is based on [1] Equation 4, and we adopt their modifications to
// alphaRoughness as input as originally proposed in [2].
float geometricOcclusion(PBRInfo pbrInputs, PBRLightInfo pbrLight)
{
    float NdotL = pbrLight.NdotL;
    float NdotV = pbrLight.NdotV;
    float r = pbrInputs.alphaRoughness;

    float attenuationL = 2.0 * NdotL / (NdotL + sqrt(r * r + (1.0 - r * r) * (NdotL * NdotL)));
    float attenuationV = 2.0 * NdotV / (NdotV + sqrt(r * r + (1.0 - r * r) * (NdotV * NdotV)));
    return attenuationL * attenuationV;
}

// The following equation(s) model the distribution of microfacet normals across the area being drawn (aka D())
// Implementation from "Average Irregularity Representation of a Roughened Surface for Ray Reflection" by T. S. Trowbridge, and K. P. Reitz
// Follows the distribution function recommended in the SIGGRAPH 2013 course notes from EPIC Games [1], Equation 3.
float microfacetDistribution(PBRInfo pbrInputs, PBRLightInfo pbrLight)
{
    float roughnessSq = pbrInputs.alphaRoughness * pbrInputs.alphaRoughness;
    float f = (pbrLight.NdotH * roughnessSq - pbrLight.NdotH) * pbrLight.NdotH + 1.0;
    return roughnessSq / (M_PI * f * f);
}

vec3 pbrModel(PBRInfo pbrInputs, vec3 lightColor, vec3 lightDir) {

    vec3 n = getNormal();                             // normal at surface point
    vec3 v = normalize(CamDir);                       // Vector from surface point to camera
    vec3 l = normalize(lightDir);                     // Vector from surface point to light
    vec3 h = normalize(l+v);                          // Half vector between both l and v
    vec3 reflection = -normalize(reflect(v, n));

    float NdotL = clamp(dot(n, l), 0.001, 1.0);
    float NdotV = abs(dot(n, v)) + 0.001;
    float NdotH = clamp(dot(n, h), 0.0, 1.0);
    float LdotH = clamp(dot(l, h), 0.0, 1.0);
    float VdotH = clamp(dot(v, h), 0.0, 1.0);

    PBRLightInfo pbrLight = PBRLightInfo(
        NdotL,
        NdotV,
        NdotH,
        LdotH,
        VdotH
    );

    // Calculate the shading terms for the microfacet specular shading model
    vec3 F = specularReflection(pbrInputs, pbrLight);
    float G = geometricOcclusion(pbrInputs, pbrLight);
    float D = microfacetDistribution(pbrInputs, pbrLight);

    // Calculation of analytical lighting contribution
    vec3 diffuseContrib = (1.0 - F) * diffuse(pbrInputs);
    vec3 specContrib = F * G * D / (4.0 * NdotL * NdotV);
    // Obtain final intensity as reflectance (BRDF) scaled by the energy of the light (cosine law)
    vec3 color = NdotL * lightColor * (diffuseContrib + specContrib);

    return color;
}

void main() {

    float perceptualRoughness = uRoughnessFactor;
    float metallic = uMetallicFactor;

#ifdef HAS_METALROUGHNESSMAP
    // Roughness is stored in the 'g' channel, metallic is stored in the 'b' channel.
    // This layout intentionally reserves the 'r' channel for (optional) occlusion map data
    vec4 mrSample = texture(uMetallicRoughnessSampler, FragTexcoord);
    perceptualRoughness = mrSample.g * perceptualRoughness;
    metallic = mrSample.b * metallic;
#endif

    perceptualRoughness = clamp(perceptualRoughness, c_MinRoughness, 1.0);
    metallic = clamp(metallic, 0.0, 1.0);
    // Roughness is authored as perceptual roughness; as is convention,
    // convert to material roughness by squaring the perceptual roughness [2].
    float alphaRoughness = perceptualRoughness * perceptualRoughness;

    // The albedo may be defined from a base texture or a flat color
#ifdef HAS_BASECOLORMAP
    vec4 baseColor = SRGBtoLINEAR(texture(uBaseColorSampler, FragTexcoord)) * uBaseColor;
#else
    vec4 baseColor = uBaseColor;
#endif

    vec3 f0 = vec3(0.04);
    vec3 diffuseColor = baseColor.rgb * (vec3(1.0) - f0);
    diffuseColor *= 1.0 - metallic;

    vec3 specularColor = mix(f0, baseColor.rgb, uMetallicFactor);

    // Compute reflectance.
    float reflectance = max(max(specularColor.r, specularColor.g), specularColor.b);

    // For typical incident reflectance range (between 4% to 100%) set the grazing reflectance to 100% for typical fresnel effect.
    // For very low reflectance range on highly diffuse objects (below 4%), incrementally reduce grazing reflectance to 0%.
    float reflectance90 = clamp(reflectance * 25.0, 0.0, 1.0);
    vec3 specularEnvironmentR0 = specularColor.rgb;
    vec3 specularEnvironmentR90 = vec3(1.0, 1.0, 1.0) * reflectance90;

    PBRInfo pbrInputs = PBRInfo(
        perceptualRoughness,
        metallic,
        specularEnvironmentR0,
        specularEnvironmentR90,
        alphaRoughness,
        diffuseColor,
        specularColor
    );

//    vec3 normal = getNormal();
    vec3 color = vec3(0.0);

#if AMB_LIGHTS>0
    // Ambient lights
    for (int i = 0; i < AMB_LIGHTS; i++) {
        color += AmbientLightColor[i] * pbrInputs.diffuseColor;
    }
#endif

#if DIR_LIGHTS>0
    // Directional lights
    for (int i = 0; i < DIR_LIGHTS; i++) {
        // Diffuse reflection
        // DirLightPosition is the direction of the current light
        vec3 lightDirection = normalize(DirLightPosition(i));
        // PBR
        color += pbrModel(pbrInputs, DirLightColor(i), lightDirection);
    }
#endif

#if POINT_LIGHTS>0
    // Point lights
    for (int i = 0; i < POINT_LIGHTS; i++) {
        // Common calculations
        // Calculates the direction and distance from the current vertex to this point light.
        vec3 lightDirection = PointLightPosition(i) - vec3(Position);
        float lightDistance = length(lightDirection);
        // Normalizes the lightDirection
        lightDirection = lightDirection / lightDistance;
        // Calculates the attenuation due to the distance of the light
        float attenuation = 1.0 / (1.0 + PointLightLinearDecay(i) * lightDistance +
            PointLightQuadraticDecay(i) * lightDistance * lightDistance);
        vec3 attenuatedColor = PointLightColor(i) * attenuation;
        // PBR
        color += pbrModel(pbrInputs, attenuatedColor, lightDirection);
    }
#endif

#if SPOT_LIGHTS>0
    for (int i = 0; i < SPOT_LIGHTS; i++) {

        // Calculates the direction and distance from the current vertex to this spot light.
        vec3 lightDirection = SpotLightPosition(i) - vec3(Position);
        float lightDistance = length(lightDirection);
        lightDirection = lightDirection / lightDistance;

        // Calculates the attenuation due to the distance of the light
        float attenuation = 1.0 / (1.0 + SpotLightLinearDecay(i) * lightDistance +
            SpotLightQuadraticDecay(i) * lightDistance * lightDistance);

        // Calculates the angle between the vertex direction and spot direction
        // If this angle is greater than the cutoff the spotlight will not contribute
        // to the final color.
        float angle = acos(dot(-lightDirection, SpotLightDirection(i)));
        float cutoff = radians(clamp(SpotLightCutoffAngle(i), 0.0, 90.0));

        if (angle < cutoff) {
            float spotFactor = pow(dot(-lightDirection, SpotLightDirection(i)), SpotLightAngularDecay(i));
            vec3 attenuatedColor = SpotLightColor(i) * attenuation * spotFactor;
            // PBR
            color += pbrModel(pbrInputs, attenuatedColor, lightDirection);
        }
    }
#endif

    // Calculate lighting contribution from image based lighting source (IBL)
//#ifdef USE_IBL
//    color += getIBLContribution(pbrInputs, n, reflection);
//#endif

    // Apply optional PBR terms for additional (optional) shading
#ifdef HAS_OCCLUSIONMAP
    float ao = texture(uOcclusionSampler, FragTexcoord).r;
    color = mix(color, color * ao, 1.0);//, uOcclusionStrength);
#endif

#ifdef HAS_EMISSIVEMAP
    vec3 emissive = SRGBtoLINEAR(texture(uEmissiveSampler, FragTexcoord)).rgb * vec3(uEmissiveColor);
#else
    vec3 emissive = vec3(uEmissiveColor);
#endif
    color += emissive;

    // Base Color
//    FragColor = baseColor;

    // Normal
//    FragColor = vec4(n, 1.0);

    // Emissive Color
//    FragColor = vec4(emissive, 1.0);

    // F
//    color = F;

    // G
//    color = vec3(G);

    // D
//    color = vec3(D);

    // Specular
//    color = specContrib;

    // Diffuse
//    color = diffuseContrib;

    // Roughness
//    color = vec3(perceptualRoughness);

    // Metallic
//    color = vec3(metallic);

    // Final fragment color
    FragColor = vec4(pow(color,vec3(1.0/2.2)), baseColor.a);
}
`

const physical_vertex_source = `//
// Physically Based Shading of a microfacet surface material - Vertex Shader
// Modified from reference implementation at https://github.com/KhronosGroup/glTF-WebGL-PBR
//
#include <attributes>

// Model uniforms
uniform mat4 ModelViewMatrix;
uniform mat3 NormalMatrix;
uniform mat4 MVP;

#include <morphtarget_vertex_declaration>
#include <bones_vertex_declaration>

// Output variables for Fragment shader
out vec3 Position;
out vec3 Normal;
out vec3 CamDir;
out vec2 FragTexcoord;

void main() {

    // Transform this vertex position to camera coordinates.
    Position = vec3(ModelViewMatrix * vec4(VertexPosition, 1.0));

    // Transform this vertex normal to camera coordinates.
    Normal = normalize(NormalMatrix * VertexNormal);

    // Calculate the direction vector from the vertex to the camera
    // The camera is at 0,0,0
    CamDir = normalize(-Position.xyz);

    // Output texture coordinates to fragment shader
    FragTexcoord = VertexTexcoord;

    vec3 vPosition = VertexPosition;
    mat4 finalWorld = mat4(1.0);
    #include <morphtarget_vertex>
    #include <bones_vertex>

    gl_Position = MVP * finalWorld * vec4(vPosition, 1.0);

}
`

const point_fragment_source = `precision highp float;

#include <material>

// GLSL 3.30 does not allow indexing texture sampler with non constant values.
// This macro is used to mix the texture with the specified index with the material color.
// It should be called for each texture index.
#if MAT_TEXTURES > 0
vec4 MIX_POINT_TEXTURE(vec4 texMixed, mat2 rotation, int i) {                                                           \
    if (MatTexVisible(i)) {                                                                                      \
        vec2 pt = gl_PointCoord - vec2(0.5);                                                                     \
        vec4 texColor = texture(MatTexture[i], (rotation * pt + vec2(0.5)) * MatTexRepeat(i) + MatTexOffset(i)); \
        if (i == 0) {                                                                                            \
            texMixed = texColor;                                                                                 \
        } else {                                                                                                 \
            texMixed = mix(texMixed, texColor, texColor.a);                                                      \
        }                                                                                                        \
    }
    return texMixed;
}
#endif

// Inputs from vertex shader
in vec3 Color;
flat in mat2 Rotation;

// Output
out vec4 FragColor;

void main() {

    // Mix material color with textures colors
    vec4 texMixed = vec4(1);
    #if MAT_TEXTURES==1
        texMixed = MIX_POINT_TEXTURE(texMixed, Rotation, 0);
    #elif MAT_TEXTURES==2
        texMixed = MIX_POINT_TEXTURE(texMixed, Rotation, 0);
        texMixed = MIX_POINT_TEXTURE(texMixed, Rotation, 1);
    #elif MAT_TEXTURES==3
        texMixed = MIX_POINT_TEXTURE(texMixed, Rotation, 0);
        texMixed = MIX_POINT_TEXTURE(texMixed, Rotation, 1);
        texMixed = MIX_POINT_TEXTURE(texMixed, Rotation, 2);
    #endif

    // Generates final color
    FragColor = min(vec4(Color, MatOpacity) * texMixed, vec4(1));
}
`

const point_vertex_source = `#include <attributes>

// Model uniforms
uniform mat4 MVP;
uniform mat4 MV;

// Material uniforms
#include <material>

// Outputs for fragment shader
out vec3 Color;
flat out mat2 Rotation;

void main() {

    // Rotation matrix for fragment shader
    float rotSin = sin(MatPointRotationZ);
    float rotCos = cos(MatPointRotationZ);
    Rotation = mat2(rotCos, rotSin, - rotSin, rotCos);

    // Sets the vertex position
    vec4 pos = MVP * vec4(VertexPosition, 1.0);
    gl_Position = pos;

    // Sets the size of the rasterized point decreasing with distance
    vec4 posMV = MV * vec4(VertexPosition, 1.0);
    gl_PointSize = MatPointSize / -posMV.z;

    // Outputs color
    Color = MatEmissiveColor;
}

`

const standard_fragment_source = `precision highp float;

// Inputs from vertex shader
in vec4 Position;     // Fragment position in camera coordinates
in vec3 Normal;       // Fragment normal in camera coordinates
in vec2 FragTexcoord; // Fragment texture coordinates

#include <lights>
#include <material>
#include <phong_model>

// Final fragment color
out vec4 FragColor;

void main() {

    // Compute final texture color
    vec4 texMixed = vec4(1);
    #if MAT_TEXTURES > 0
        bool firstTex = true;
        if (MatTexVisible(0)) {
            vec4 texColor = texture(MatTexture[0], FragTexcoord * MatTexRepeat(0) + MatTexOffset(0));
            if (firstTex) {
                texMixed = texColor;
                firstTex = false;
            } else {
                texMixed = Blend(texMixed, texColor);
            }
        }
        #if MAT_TEXTURES > 1
            if (MatTexVisible(1)) {
                vec4 texColor = texture(MatTexture[1], FragTexcoord * MatTexRepeat(1) + MatTexOffset(1));
                if (firstTex) {
                    texMixed = texColor;
                    firstTex = false;
                } else {
                    texMixed = Blend(texMixed, texColor);
                }
            }
            #if MAT_TEXTURES > 2
                if (MatTexVisible(2)) {
                    vec4 texColor = texture(MatTexture[2], FragTexcoord * MatTexRepeat(2) + MatTexOffset(2));
                    if (firstTex) {
                        texMixed = texColor;
                        firstTex = false;
                    } else {
                        texMixed = Blend(texMixed, texColor);
                    }
                }
            #endif
        #endif
    #endif

    // Combine material with texture colors
    vec4 matDiffuse = vec4(MatDiffuseColor, MatOpacity) * texMixed;
    vec4 matAmbient = vec4(MatAmbientColor, MatOpacity) * texMixed;

    // Normalize interpolated normal as it may have shrinked
    vec3 fragNormal = normalize(Normal);

    // Calculate the direction vector from the fragment to the camera (origin)
    vec3 camDir = normalize(-Position.xyz);

    // Workaround for gl_FrontFacing
    vec3 fdx = dFdx(Position.xyz);
    vec3 fdy = dFdy(Position.xyz);
    vec3 faceNormal = normalize(cross(fdx,fdy));
    if (dot(fragNormal, faceNormal) < 0.0) { // Back-facing
        fragNormal = -fragNormal;
    }

    // Calculates the Ambient+Diffuse and Specular colors for this fragment using the Phong model.
    vec3 Ambdiff, Spec;
    phongModel(Position, fragNormal, camDir, vec3(matAmbient), vec3(matDiffuse), Ambdiff, Spec);

    // Final fragment color
    FragColor = min(vec4(Ambdiff + Spec, matDiffuse.a), vec4(1.0));
}
`

const standard_vertex_source = `#include <attributes>

// Model uniforms
uniform mat4 ModelViewMatrix;
uniform mat3 NormalMatrix;
uniform mat4 MVP;

#include <material>
#include <morphtarget_vertex_declaration>
#include <bones_vertex_declaration>

// Output variables for Fragment shader
out vec4 Position;
out vec3 Normal;
out vec2 FragTexcoord;

void main() {

    // Transform vertex position to camera coordinates
    Position = ModelViewMatrix * vec4(VertexPosition, 1.0);

    // Transform vertex normal to camera coordinates
    Normal = normalize(NormalMatrix * VertexNormal);

    vec2 texcoord = VertexTexcoord;
#if MAT_TEXTURES > 0
    // Flip texture coordinate Y if requested.
    if (MatTexFlipY(0)) {
        texcoord.y = 1.0 - texcoord.y;
    }
#endif
    FragTexcoord = texcoord;
    vec3 vPosition = VertexPosition;
    mat4 finalWorld = mat4(1.0);
    #include <morphtarget_vertex>
    #include <bones_vertex>

    // Output projected and transformed vertex position
    gl_Position = MVP * finalWorld * vec4(vPosition, 1.0);
}
`

// Maps include name with its source code
var includeMap = map[string]string{

	"attributes":                      include_attributes_source,
	"bones_vertex":                    include_bones_vertex_source,
	"bones_vertex_declaration":        include_bones_vertex_declaration_source,
	"lights":                          include_lights_source,
	"material":                        include_material_source,
	"morphtarget_vertex":              include_morphtarget_vertex_source,
	"morphtarget_vertex2":             include_morphtarget_vertex2_source,
	"morphtarget_vertex_declaration":  include_morphtarget_vertex_declaration_source,
	"morphtarget_vertex_declaration2": include_morphtarget_vertex_declaration2_source,
	"phong_model":                     include_phong_model_source,
}

// Maps shader name with its source code
var shaderMap = map[string]string{

	"basic_fragment":    basic_fragment_source,
	"basic_vertex":      basic_vertex_source,
	"panel_fragment":    panel_fragment_source,
	"panel_vertex":      panel_vertex_source,
	"physical_fragment": physical_fragment_source,
	"physical_vertex":   physical_vertex_source,
	"point_fragment":    point_fragment_source,
	"point_vertex":      point_vertex_source,
	"standard_fragment": standard_fragment_source,
	"standard_vertex":   standard_vertex_source,
}

// Maps program name with Proginfo struct with shaders names
var programMap = map[string]ProgramInfo{

	"basic":    {"basic_vertex", "basic_fragment", ""},
	"panel":    {"panel_vertex", "panel_fragment", ""},
	"physical": {"physical_vertex", "physical_fragment", ""},
	"point":    {"point_vertex", "point_fragment", ""},
	"standard": {"standard_vertex", "standard_fragment", ""},
}
